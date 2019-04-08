# 04_tryto_sublevel
#the output is 
---------------------------------- 
get by key :- 
---------------------------------- 
{"Id":"0","Name":"ali","Group":["A","B"]}
---------------------------------- 
get all data in database :- 
---------------------------------- 
your key :- 0-ali || your data :- {"Id":"0","Name":"ali","Group":["A","B"]}
your key :- 1-ola || your data :- {"Id":"1","Name":"ola","Group":["B","c"]}
your key :- 2-aya || your data :- {"Id":"2","Name":"aya","Group":["C","D"]}
your key :- 3-mi || your data :- {"Id":"3","Name":"mi","Group":["D","E"]}
your key :- 4-noha || your data :- {"Id":"4","Name":"noha","Group":["E","F"]}
your key :- 5-nour || your data :- {"Id":"5","Name":"nour","Group":["F","G"]}
---------------------------------- 
search on database :- try to find the name aya 
---------------------------------- 
the result is founded
your key :- 2-aya || your data :- {"Id":"2","Name":"aya","Group":["C","D"]}
---------------------------------- 
get all data in database by id :- 
---------------------------------- 
your key :- 0-ali || your data :- {"Id":"0","Name":"ali","Group":["A","B"]}
---------------------------------- 
your fist key :- 
---------------------------------- 
your key :- 0-ali || your data :- {"Id":"0","Name":"ali","Group":["A","B"]} || is first :true
---------------------------------- 
your last key :- 
---------------------------------- 
your key :- 5-nour || your data :- {"Id":"5","Name":"nour","Group":["F","G"]} || is Last :true
---------------------------------- 
update by key  :- update <0-ali> key put him in group A , B and C
---------------------------------- 
get by key :- {"Id":"0","Name":"ali","Group":["A","B","C"]}
---------------------------------- 
delete by key  :- delete <1-ola> 
---------------------------------- 
deleted >>>
find 1-ola key ... 
err.. Not founded
2019/04/08 22:22:04 leveldb: not found
exit status 1