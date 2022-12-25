# Backuper Golang 

It's a simple tool to do kind of backups and copyings of folder.
You can create point which will represent name, source dist. and output dist. and reuse it in one command.

1. ## Installation
 - Download backuper binary file
 - Put it in /usr/local/bin
 - run with sudo
2. ## Arguments
 - `backuper init` - creates a config file in /etc/backuper where your points will locate
 - `backuper add <name> <source_dir> <output_dir>` - add point to config for reuse
 - `backuper list` - print a list of all points
 - `backuper remove <name>` - remove point
 - `backuper exec <name>` 
    ### Flags
    - `-m` does a merge of your src and output folders
    - `-r` remove old backuped folder (if exists) and create new
    - _noflag_ create folder with _unixtime milliseconds now

    Features to add
    - `backuper edit` - to edit point
    - `backuper add -i <hours>` - to create interval which specify how often backups will be created
    - deamon for non stop running with "nextBackup: time.Date"