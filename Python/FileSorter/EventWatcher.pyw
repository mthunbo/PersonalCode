import time

from pathlib import Path
from watchdog.observers import Observer
from EventSorter import SorterEventHandler

class FileWatcher:
    def __init__(self, src_path):
        self.__src_path = src_path
        self.__event_handler = SorterEventHandler()
        self.__event_observer = Observer()

    def run(self):
        self.start()
        try:
            while True:
                time.sleep(10)
        except KeyboardInterrupt:
            self.stop()

    def start(self):
        self.__schedule()
        self.__event_observer.start()

    def stop(self):
        self.__event_observer.stop()
        self.__event_observer.join()

    def __schedule(self):
        self.__event_observer.schedule(
            self.__event_handler,
            self.__src_path,
            recursive=True
        )

if __name__ == "__main__":
    src_path = Path(f'{str(Path.home())}/Downloads')
    FileWatcher(src_path).run() 
