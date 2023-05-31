var fs = require('fs');

fs.appendFile('HelloSeaWorld.html', 'This is my text', function (err){
    if (err) throw err;
    console.log('Saved!')
})