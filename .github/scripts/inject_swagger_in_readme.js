const fs = require('fs');
const { exit } = require('process');



function readSwaggerLines() {
    if (!fs.existsSync("SWAGGER.md")) {
        console.error("Missing SWAGGER.md")
        exit(1)
    }
    const allFileContents = fs.readFileSync("SWAGGER.md", "utf-8");
    let ignore = false
    const newFileContents = []
    allFileContents.split(/\r?\n/).forEach(line => {
        line = line.trimEnd();
        if (line == '---') {
            ignore = !ignore
            return
        }
        if (!ignore) {
            newFileContents.push(line)
        }
    })
    newFileContents.push(`_swagger data generated @ ${new Date()}_`, '')
    return newFileContents
}
function getReadmeLines(newSwaggerLines) {
    if (!fs.existsSync('README.md')) {
        console.error('Missing README.md')
        exit(1)
    }
    const allFileContents = fs.readFileSync('README.md', 'utf-8');
    const newFileContents = []
    let startedSwagger = false

    allFileContents.split(/\r?\n/).forEach(line => {
        if (line == "## Swagger documentation") {
            newFileContents.push(line, '', ...newSwaggerLines)
            startedSwagger = true
        } else {
            if (line == "* end documentation") {
                startedSwagger = false
            }
            if (!startedSwagger) {
                newFileContents.push(line)
            }
        }
    });
    return newFileContents
}
console.log("Updating swagger in README.md")

let newSwaggerLines = readSwaggerLines()
let newReadmeLines = getReadmeLines(newSwaggerLines)
fs.writeFileSync("README.md", newReadmeLines.join("\n"))


const used = process.memoryUsage().heapUsed / 1024 / 1024;
console.log(`The script uses approximately ${Math.round(used * 100) / 100} MB`);
