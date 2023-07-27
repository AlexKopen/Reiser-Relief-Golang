module.exports = {
    plugins: [
        require('postcss-hash')({
            manifest: './static/manifest.json',
            hashFunction: 'md5',
            hashDigest: 'hex',
            trim: 20,
        })
    ]
};
