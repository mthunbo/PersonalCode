from pathlib import Path
from os import scandir, rename
from time import sleep
from shutil import move, Error
from watchdog.events import FileSystemEventHandler
from os.path import splitext, exists, join

import json

# Class that can go through the files in a given folder and sort them by category given in a json file
class SorterEventHandler(FileSystemEventHandler):

    src_path = ''
    with open('config.json', encoding='utf-8') as conf:
        cats = json.load(conf)
    extensionMap = {}

    def __init__(self):
        self.src_path = Path(f'{str(Path.home())}/Downloads')
        for cat in self.cats:
            folderName = cat['name']
            folderPath = Path(f'{self.src_path}/{folderName}')
            if not folderPath.exists():
                folderPath.mkdir(parents=True, exist_ok=True)
            for extention in cat['extensions']:
                self.extensionMap[extention] = folderName        
    
    def on_modified(self, event):
        self.process(event)

    def process(self, event):
        with scandir(self.src_path) as entries:
            for entry in entries:
                if entry.is_file() and not entry.name.startswith('.'):
                    name = entry.name
                    destFolder = Path(f'{self.src_path}/{self.check_extension(entry, name)}')
                    self.moveFile(destFolder, entry, name)


    # Creates a unique name for the given file if a file of the same name is found in the destination folder
    def uniqueName(self, dest, name):
        filename, extension = splitext(name)
        counter = 1
        while exists(f"{dest}/{name}"):
            name = f"{filename}({str(counter)}){extension}"
            counter += 1
        return name

    # Moves the given file to it's destination folder after checking if the name is unique
    def moveFile(self, dest, entry, name):
        try:
            if exists(f"{dest}/{name}"):
                unique_name = self.uniqueName(dest, name)
                oldName = join(dest, name)
                newName = join(dest, unique_name)
                rename(oldName, newName)
            move(entry, dest)
        except Error as e:
            print(e)

    # Checks the extension of the given file and returns the category it belongs to. Category is determined på config.json
    def check_extension(self, entry, name):
        fileName, fileExt = splitext(name)
        
        for cat in self.cats:
            currentFileType = cat['name']
            for extensions in cat['extensions']:
                if fileExt == extensions or fileExt == extensions.upper():
                    return currentFileType
        return "Others"
            