//Sample code to read in test cases:
var fs  = require("fs");

fs.readFileSync(process.argv[2]).toString().split('\n').forEach(function (line) {
    if (line != "") {
        //do something here
        var words = line.split(' ');
        var longestWord = '';
        for (var i = 0; i < words.length - 1; i++) {
            if (words[i].length > longestWord.length) {
                longestWord = words[i];
            }
        }
        console.log(longestWord);
    }
});
