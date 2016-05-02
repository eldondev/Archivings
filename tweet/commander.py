import json
keys = json.load(open('keys.json'))
sha_key = "%s&%s" % (keys['consumer_secret'], keys['access_token_secret'])

from hashlib import sha1
import hmac
requests = json.load(open('config.json'))
method = requests['method']
del requests['method']
import urllib.request, urllib.error, urllib.parse
url = requests['url']
del requests['url']
import codecs
hencode = codecs.getencoder('hex_codec')
aencode = codecs.getdecoder('utf_8')
requests['oauth_nonce']=aencode(hencode(open('/dev/urandom', 'rb').read(16))[0])[0]
import time
requests['oauth_timestamp']=str(int(time.time()))
requests = list(requests.items())
requests.sort()
oauth_string = urllib.parse.quote('&'.join(['%s=%s' % (k,v) for k,v in requests]),'')
request_string = '&'.join([method,urllib.parse.quote(url,''),oauth_string])

base64encode = codecs.getencoder('base64_codec')
import hmac
import urllib.parse
requests.append(('oauth_signature', urllib.parse.quote(base64encode(hmac.new(bytes(sha_key, 'utf8'),bytes(request_string, 'utf8'),sha1).digest())[0][:-1])))
requests.sort()
header = str('Authorization: OAuth ' + ', '.join(['%s="%s"' % k for k in requests]))

print(url)
print("'-H%s'" % header)

