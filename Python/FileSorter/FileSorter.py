import shutil
import json

from pathlib import Path

class FileEventHandler():
    # Method that takes a file to move and a destination folder as arguements, checks for destination folder if it isn't there it creates it and moves the file there
    def fileMover(file, folder):
        try:
            if not folder.exists():
                folder.mkdir(parents=True, exist_ok=True)
            shutil.move(file, folder)
        except shutil.Error as e:
            print(e)

    # Method that sorts through the files in the folder given as an arguement
    def folderSorter(path):
        with open('config.json', encoding='utf-8') as conf:
            cats = json.load(conf)

        #Creates a map with the names and extensions of the files to be able to sort them
        extentionMap = {}
        for cat in cats:
            folderName = cat['name']
            for extention in cat['extensions']:
                extentionMap[extention] = folderName

        for file in path.iterdir():
            if file.is_file() and not file.name.startswith('.'):
                destFolder = extentionMap.get(file.suffix, 'Other')
                FileEventHandler.fileMover(file, file.parent.joinpath(destFolder))

if __name__ == '__main__':
    downloadFolder = Path(f'{str(Path.home())}/Downloads')
    FileEventHandler.folderSorter(downloadFolder)