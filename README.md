# nginZ
this tool is based on this research : <br />
Part one https://blog.detectify.com/2020/11/10/common-nginx-misconfigurations/ <br />
Part tow https://labs.detectify.com/2021/02/18/middleware-middleware-everywhere-and-lots-of-misconfigurations-to-fix/<br />
Big thank to detectify community  https://detectify.com<br />
this tool was created in collaboration with my friend https://twitter.com/medmahmoudi_619 
# what is nginZ
nginZ is scanner for common Nginx misconfigurations and vulnerabilities
# vulnerabilities that can be detected
Off-By-Slash<br />
Unsafe variable use ( XSS by SCRIPT_NAME) (CRLF Injection by $uri)<br />
nginx variable reflected in response<br />
HTTP Request Splitting with cloud storage<br />
Controlling proxied host<br />
# install
`go install -v github.com/raoufmaklouf/nginZ@latest`
# Usage
`cat nginx_Urls.txt | nginZ` <br /><br /><br />


![alt text](https://github.com/raoufmaklouf/nginZ/blob/main/nginZ.jpeg)

