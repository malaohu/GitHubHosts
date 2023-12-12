import http.client

conn = http.client.HTTPSConnection("ipaddress.com")
payload = ''
headers = {
  'Host': 'ipaddress.com',
  'Cookie': 'active_template::280870=pub_site.1653574012; ezepvv=2352; ezoab_280870=mod1-c; ezoadgid_280870=-1; ezopvc_280870=3; ezoref_280870=; ezosuibasgeneris-1=e77ba9d1-e8dd-405b-585b-1eb55ad41936; ezovid_280870=885070763; ezovuuid_280870=5afdf2f2-970d-4d09-4b32-7ceda11e00ac; ezovuuidtime_280870=1653574013; lp_280870=https://ipaddress.com/website/github.com'
}
conn.request("GET", "/website/alive.github.com", payload, headers)
res = conn.getresponse()
data = res.read()
print(data.decode("utf-8"))
