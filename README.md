# 04_tryto_sublevel
#the ount but is 
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
your key :- 0-ali || your data :- {"Id":"0","Name":"ali","Group":["A","B"]} || is Last :true
---------------------------------- 
update by key  :- update <0-ali> key put him in group F and A 
---------------------------------- 
get by key :- {"Id":"0","Name":"ali","Group":["F","A"]}
---------------------------------- 
delete by key  :- delete <1-ola> 
---------------------------------- 
deleted >>>
find 1-ola key ... 
err.. Not founded
2019/04/08 17:35:54 leveldb: not found
exit status 1