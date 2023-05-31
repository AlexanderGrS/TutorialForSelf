var fs = require('fs')

fs.writeFile('HelloSeaWorld.txt', 'This is my text',function(err){
    if (err) throw err;
    console.log('Saved!')
})
