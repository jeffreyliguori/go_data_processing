This is some elections-related go stuff I'm working on because I'm bored.

I took the data from https://github.com/tonmcg/County_Level_Election_Results_12-16.git ; the person that assembled this deserves 99% of the credit as the code here is incredibly simple.

The data in states/ was taken from http://www.census.gov/popest/data/counties/asrh/2011/CC-EST2011-RACE6.html with this grossness:for i in $(echo "01 02 03 04 05 06 07 08 09" && for ((i=10; i<=56; i++)); do echo $i; done); do curl -s http://www.census.gov/popest/data/counties/asrh/2011/files/CC-EST2011-6RACE-$i.csv > states/$i.csv ; done , then to find the files that didn't exist I ran grep \< * | cut -d':' -f1 | uniq
