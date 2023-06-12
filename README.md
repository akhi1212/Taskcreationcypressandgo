# Taskcreationcypressandgo

Pre requisites 

ensure rancher desktop  setup  


run command 
sudo docker run --privileged -d --restart=unless-stopped -p 80:80 -p 443:443 rancher/rancher

verify  manually https://localhost/dashboard

install cypress framework 

 Run cypress 
 go to folder where your project  located and run  
                    npx cypress open
                    
                    
from cypress gui  click e2e testing select browser & click e2e in <selected browser>
  
then select file WebUIRancher.cy.js file click that file , it will start execute 
