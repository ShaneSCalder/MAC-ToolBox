const fs = require('fs');
const path = require('path');

// Correct paths based on the script's location
const crawledDirPath = path.join(__dirname, '..', 'data', 'crawled');
const indexedDirPath = path.join(__dirname, '..', 'data', 'indexed');

// Ensure the indexed directory exists
if (!fs.existsSync(indexedDirPath)) {
    fs.mkdirSync(indexedDirPath, { recursive: true });
}

// Function to safely read and parse JSON files
function safeReadAndParseJSON(filePath) {
    try {
        const fileContent = fs.readFileSync(filePath, 'utf8');
        return JSON.parse(fileContent);
    } catch (error) {
        console.error(`Error reading or parsing JSON from file ${filePath}:`, error);
        return null; // Return null if there was an error
    }
}

// Function to process and index files
function processAndIndexFiles() {
    const files = fs.readdirSync(crawledDirPath);
    const index = [];

    files.forEach(file => {
        const filePath = path.join(crawledDirPath, file);
        const data = safeReadAndParseJSON(filePath);

        // Proceed only if data could be successfully read and parsed
        if (data) {
            if (data.DCMI_Metadata) {
                // Extract relevant information for Dublin Core metadata
                const { Title, Creator, Date, Identifier } = data.DCMI_Metadata;
                index.push({ Type: 'Dublin Core', Title, Creator, Date, Identifier });
            } else if (data.Cids) {
                // Process each CID in the file
                data.Cids.forEach(cid => {
                    index.push({ Type: 'CID', Path: cid['/'] });
                });
            }
        }
    });

    // Save the index to a file in the indexed directory
    const indexPath = path.join(indexedDirPath, 'index.json');
    fs.writeFileSync(indexPath, JSON.stringify(index, null, 2));
    console.log(`Indexing complete. Indexed data saved to ${indexPath}`);
}

// Execute the indexing process
processAndIndexFiles();
