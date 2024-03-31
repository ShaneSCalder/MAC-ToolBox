const fs = require('fs');
const path = require('path');

// Path to the indexed data file
const indexPath = path.join(__dirname, '..', 'data', 'indexed', 'index.json');

// Function to load indexed data
function loadIndexedData() {
    try {
        const data = fs.readFileSync(indexPath, 'utf8');
        return JSON.parse(data);
    } catch (error) {
        console.error(`Failed to read or parse index file: ${error.message}`);
        process.exit(1);
    }
}

// Simple search function
function search(query) {
    const index = loadIndexedData();
    // Assuming each indexed item has a 'Title' property to search against
    // Adjust based on your actual indexed data structure
    const results = index.filter(item => item.Title && item.Title.toLowerCase().includes(query.toLowerCase()));
    return results;
}

// Get the search query from command-line arguments
const query = process.argv[2];

// Perform the search and print the results
if (query) {
    const results = search(query);
    console.log(JSON.stringify(results));
} else {
    console.log("No search query provided.");
}
