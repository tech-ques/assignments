# Assignment

Build a sample Go app where you have a list of users, each with properties like Name, email, mobile, etc. 
Each user also has their preferred way of getting notification from app for. e.g. via email, SMS, call, etc.
 
Implement dummy methods for each of these notifications and demonstrate how they will be called for each user based on their choice of mode of notification.
 
You can leave out actual details about implementation of notifications via email, SMS, call.
Remember that focus is to be given on - 
 
1. selecting Go data type(s) to implement different notification types/pattern.
2. making use of Go-routines and appropriate communication mechanism to ensure no more than ‘10’ users are being processed at a time.
3. handling retry mechanism for any failure