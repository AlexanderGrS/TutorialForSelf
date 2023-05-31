var fs = require('fs')

fs.unlink('HelloSeaWorld.txt', function(err){
    if (err) throw err;
    console.log('Deleted!')
})