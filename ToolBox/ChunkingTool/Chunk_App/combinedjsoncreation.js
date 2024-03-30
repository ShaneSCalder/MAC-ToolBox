const fs = require('fs');
const path = require('path');

const generateOriginalMetadata = () => {
  const metadata = {
    "DCMI_Metadata": {
      "Title": "Sample Dataset of Apple Fruit Measurements",
      "Creator": "Merkle Meta Research Team",
      "Subject": "Apples, Agriculture, Fruit Measurements",
      "Description": "This dataset provides detailed measurements and quality assessments of apple fruits, collected as part of an agricultural study by Merkle Meta.",
      "Publisher": "Merkle Meta Data Publishing",
      "Contributors": ["Jane Appleseed", "John Orchard"],
      "Date": "2024-03-15",
      "Type": "Dataset",
      "Format": "CSV",
      "Identifier": "MM2024-Apple-Measurements",
      "Source": "Field Study",
      "Language": "EN",
      "Relation": "Part of the Merkle Meta Agriculture Studies Collection",
      "Coverage": "Global",
      "Rights": "Creative Commons Attribution 4.0 International (CC BY 4.0)",
      "DatasetDivisionDescription": "The dataset was divided into four equal parts, with each chunk representing approximately 25% of the total rows from the original dataset. Each chunk was created to facilitate easier analysis and handling of the data."
    }
  };

  fs.writeFileSync(path.join('CHUNK', 'JSONdata', 'inputmetadata.json'), JSON.stringify(metadata, null, 2));
};

const generateChunkMetadata = () => {
  const chunkDir = path.join('dataout');
  const jsonDir = path.join('JSONdata');
  if (!fs.existsSync(jsonDir)) {
    fs.mkdirSync(jsonDir, { recursive: true });
  }

  const chunkFiles = fs.readdirSync(chunkDir).filter(file => file.endsWith('.csv'));

  chunkFiles.forEach((file, index) => {
    const partNumber = index + 1; // Assuming index starts at 0 for the first chunk
    const totalChunks = chunkFiles.length; // Total number of chunks created

    const chunkMetadata = {
        "Title": `Sample Dataset of Apple Fruit Measurements - Chunk ${partNumber}`,
        "Identifier": `MM2024-Apple-Measurements-Chunk${partNumber}`,
        "Date": "2024-03-15",
        "Format": "CSV",
        "Relation": `This chunk is part ${partNumber} of ${totalChunks} of the MM2024-Apple-Measurements dataset.`,
        "Type": "Dataset",
        "Language": "EN",
        "Coverage": "Global",
        "Rights": "Creative Commons Attribution 4.0 International (CC BY 4.0)"
    };

    fs.writeFileSync(path.join(jsonDir, `${file}.json`), JSON.stringify(chunkMetadata, null, 2));
  });
};

const combineAllMetadata = () => {
  const jsonDir = path.join('JSONdata');
  const metadataFiles = fs.readdirSync(jsonDir).filter(file => file.endsWith('.json') && file !== 'inputjson.json');

  const originalMetadataPath = path.join(jsonDir, 'inputjson.json');
  const originalMetadata = JSON.parse(fs.readFileSync(originalMetadataPath, 'utf8'));

  originalMetadata.DCMI_Metadata.Chunks = metadataFiles.map(file => {
    return JSON.parse(fs.readFileSync(path.join(jsonDir, file), 'utf8'));
  });

  const combinedMetadataPath = path.join(jsonDir, 'combinedMetadata.json');
  fs.writeFileSync(combinedMetadataPath, JSON.stringify(originalMetadata, null, 2));

  console.log('Combined metadata has been generated.');
};

// Run the metadata generation and combination process
generateOriginalMetadata();
generateChunkMetadata();
combineAllMetadata();
