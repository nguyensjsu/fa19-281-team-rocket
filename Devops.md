# Goal
To implement Continuos Deployment for Frontend Application deployed in Heroku.

# Implementation
Heroku allows github integration but to do that we would require admin access to the repositiry. 
Github Actions allows running scripts and commands on push to the repository. This can be done without admin access.
So to implement continuos deployment, we can write a github action which pushes the front end code to heroku.

Front end code is located in git subtree src/grubhub-fe

This code can be pushed via command
````
git subtree push --prefix src/grubhub-fe https://heroku:$HEROKU_API_TOKEN@git.heroku.com/$HEROKU_APP_NAME.git master
````
Storing API token in env is unsafe and should have been done as Secret env in github. But this also requires admin access.

# Result
![Image](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/journal/CI_CD.jpg)
