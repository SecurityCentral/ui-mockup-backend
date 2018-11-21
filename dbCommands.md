Mongo DB useful commands


Show Databases
--------------
show dbs

Select myDb
------------
use myDb

Show Available Collections
--------------------------
db.getCollectionNames()


Show All user entries
---------------------
db.user.find()

Show All standard entries
---------------------
db.std.find()


Show specifc control under any standard
---------------------------------------
db.std.find({"controls.controlname": "IA-2 (1)"}).pretty();

SAMPLE API calls
----------------

Get certification - http://localhost:1377/standard/get_certification/FedRAMP%20Low
Get standard - http://localhost:1377/standard/get_standard/nist-800-53-latest
Load(GET) standard - http://localhost:1377/standard/load_standards
Load(GET) certification - http://localhost:1377/standard/load_certifications

POST for login - http://localhost:1377/user/login
PUT create user - http://localhost:1377/user/

