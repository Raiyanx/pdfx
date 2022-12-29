const dotenv = require('dotenv');
const path = require('node:path');
const fs = require('fs')

dotenv.config();

const ILovePDFApi = require('@ilovepdf/ilovepdf-nodejs');
const ILovePDFFile = require('@ilovepdf/ilovepdf-nodejs/ILovePDFFile');

const instance = new ILovePDFApi(process.env.PUBLIC_KEY, process.env.SECRET_KEY);

// Public and secret key can be found in your developer panel
// at https://developer.ilovepdf.com/user/projects .
const task = instance.newTask('merge');

// Promise-based way to use ILovePDFApi.
task.start()
.then(() => {
    const file = new ILovePDFFile(path.resolve(__dirname, "./q1.pdf")) 
    return task.addFile(file);
})
.then(() => {
    const file = new ILovePDFFile(path.resolve(__dirname, "./q2.pdf")) 
    return task.addFile(file);
})
.then(() => {
    return task.process();
})
.then(() => {
    return task.download();
})
.then((data) => {
    fs.writeFileSync('newfile.pdf', data)
    console.log('DONE');
});

