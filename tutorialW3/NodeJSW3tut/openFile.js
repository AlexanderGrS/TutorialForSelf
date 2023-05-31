var fs = require('fs')

fs.open('HelloSeaWorld.txt', 'w', function(err, file){
    if (err) throw err
    console.log('Saved!')
});