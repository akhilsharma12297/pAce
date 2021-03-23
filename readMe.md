# **pAce**

The assigment is written in Golang using Gin Gonic Framework.

As documented in the email , the service supports 4 APIs with base url as. `http://localhost:8080/`

     PUT  /transactionservice/transaction/$transaction_id
     GET  /transactionservice/transaction/$transaction_id 
     GET  /transactionservice/types/$types_value 
     GET  /transactionservice/sum/$transaction_id


[Link to Postman API collection for the service.](https://www.getpostman.com/collections/211ba44234f135969965)

**All the APIs are test and working .**

**While running please ensure to set DB_HOST , DB_DATABASE , DB_PORT , DB_USER , DB_PASSWORD @ `pAce/src/service/transactionStructs.go [11- 15]`**

Below given is the Table structure used for the transaction.

    create table POCKET_ACES_TRANSACTION
    (
	    ID INT not null,
	    AMOUNT DOUBLE not null,
	    TYPE_VALUE VARCHAR(256) not null,
	    PARENT_ID int null
    );
    
    create unique index POCKET_ACES_TRANSACTION_ID_Uindex
    on POCKET_ACES_TRANSACTION (ID);
    
    alter table POCKET_ACES_TRANSACTION
    add constraint POCKET_ACES_TRANSACTION_pk
    primary key (ID);

****************