from time import sleep
from notifypy import Notify

FILENAME="chat.txt"

def main():
    with open(FILENAME) as f:
        notification = Notify()
        for line in f.readlines():
            if line !=  "":
                notification.title = "Cool Title"
                notification.message = line
                sleep(1)
                notification.send()
            
            
main()