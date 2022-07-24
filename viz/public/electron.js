const path = require('path');

const {app, BrowserWindow} = require('electron');
const isDev = false;

const exec = require('child_process').exec;
const cueDir = app.commandLine.getSwitchValue('cueDir');

function createWindow() {
    exec(`../go/dagviz -c ${cueDir} -j=true > ./src/data/data.json`, (err, stdout, stderr) => {
        if (err) {
            console.log(stdout);
            return;
        }
        console.log(stdout);
    });
    // Create the browser window.
    const win = new BrowserWindow({
        width: 800,
        height: 600,
        webPreferences: {
            nodeIntegration: true,
        },
    });

    // and load the index.html of the app.
    // win.loadFile("index.html");
    win.loadURL(
        'http://localhost:3000'
    );
    // Open the DevTools.
    if (isDev) {
        win.webContents.openDevTools({mode: 'detach'});
    }
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.whenReady().then(createWindow);

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit();
    }
});

app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
        createWindow();
    }
});