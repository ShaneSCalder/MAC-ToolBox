const fs = require('fs');
const crypto = require('crypto');
const path = require('path');

// Function to compute the SHA-256 hash of file content
function computeHash(filePath) {
    return new Promise((resolve, reject) => {
        const hash = crypto.createHash('sha256');
        const stream = fs.createReadStream(filePath);

        stream.on('data', function(data) {
            hash.update(data, 'utf8');
        });

        stream.on('end', function() {
            resolve(hash.digest('hex'));
        });

        stream.on('error', reject);
    });
}

// Path to the chunk file
const chunkFilePath = path.join(__dirname, 'data', 'output', '// add your data path here - chunk you want to check');

// Assuming you have the stored hash for comparison
const storedHashForChunk = 'your_stored_hash_here';

computeHash(chunkFilePath)
    .then(computedHash => {
        console.log(`Computed Hash: ${computedHash}`);
        console.log(`Stored Hash: ${storedHashForChunk}`);

        if (computedHash === storedHashForChunk) {
            console.log('Verification successful: The chunk is intact and unaltered.');
        } else {
            console.log('Verification failed: The chunk has been altered or is not the correct chunk.');
        }
    })
    .catch(err => {
        console.error('Error computing hash:', err);
    });
