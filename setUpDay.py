  
import datetime
import os
import requests
import shutil
import sys
import argparse

def main(year, day):
    print("setting up AoC problem for {} day {}".format(year, day))
    basePath = os.path.dirname(os.path.realpath(__file__))
    todayPath = os.path.join(basePath,str(year),"day{:02d}".format(day))
    templatePath = os.path.join(basePath,"dayTemplate")

    if not os.path.exists(todayPath):
        os.makedirs(todayPath)

    print(todayPath)

    for filename in os.listdir(templatePath):
        if filename == "input":
            continue
        source = os.path.join(templatePath, filename)
        target = os.path.join(todayPath, filename)
        if not os.path.exists(os.path.join(todayPath,filename)):
            shutil.copyfile(src=source, dst=target)

    inputFile = os.path.join(todayPath,"input")
    if not os.path.exists(inputFile):
        getInput(year,day,inputFile)

def getInput(year,day,savePath):
    inputURL = "https://adventofcode.com/{}/day/{}/input".format(year,day)

    with open("session.cookie",'r') as sessionFile:
        session = sessionFile.read().replace('\n', '')
        if not session:
            sys.exit("Session not specified")

        cookies = {
            "session": session,
        }

        attemptsRemaining = 5
        while attemptsRemaining > 0:
            attemptsRemaining-=1

        result = requests.get(inputURL,cookies=cookies,stream=True)
        if result.status_code == 200:
            with open(savePath, 'wb') as inputFile:
                for chunk in result:
                   inputFile.write(chunk)
                return 
        sys.exit("Failed to fetch input: {}".format(result))


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Script to setup the advent of code problem.')
    parser.add_argument('year', type=int, help='the year of the problem you want setup')
    parser.add_argument('day', type=int, help='the day of the problem you want setup')
    
    args = parser.parse_args()
    
    if args.day > 25 or args.day < 0:
        sys.exit("Invalid day {}".format(args.day))
    now = datetime.datetime.now()
    if args.year > now.year or args.year < 2015:
        sys.exit("Invalid year {}".format(args.year))

    main(args.year, args.day)
