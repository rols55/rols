function findIP(str) {
    const regex = /(?<![.\d])(?:0|[1-9]\d{0,1}|1\d{2}|2[0-4]\d|25[0-5])(?:\.(?:0|[1-9]\d{0,1}|1\d{2}|2[0-4]\d|25[0-5])){3}(?::\d{1,5})?(?![.\d])/g
    const matches = str.match(regex);
    if (!matches) {
      return [];
    }
    
    // Remove any IPs with leading zeros
    const filteredMatches = matches.filter((ip) => {
      return !/^0\d/.test(ip.replace(/:\d{1,5}$/, ''));
    });
     const validIps = filteredMatches.filter((ip) => {
        const port = ip.split(':')[1];
        return !port || parseInt(port) <= 65535;
      });
      
      return validIps.map((ip) => ip.replace(/^0+(\d)/g, '$1'));
  }
  
const str = 'qqq http:// qqqq q qqqqq https://something.com/hello qqqqqqq qhttp://example.com/hello?you=something&something=you qq 233.123.12.234 qw w wq wqw wqw ijnjjnfapsdbjnkfsdiqw klfsdjn fs fsd https://devdocs.io/javascript/global_objects/object/fromentries njnkfsdjnk sfdjn fsp fd192.168.1.123:8080 https://devdocs.io/javascript/global_objects/regexp/@@split\nhtpp://wrong/url hello %$& wf* ][½¬ http://correct/url?correct=yes è[}£§ https://nan-academy.github.io/js-training/?page=editor#data.nested 255.256.1233.2\nssages has become an accepted part of many cultures, as happened earlier with emailing. htt://[1] This makes texting a quick and http://www.example.com/mypage.html?crcat=test&crsource=test&crkw=buy-a-loteasy way to communicate 255.256.2 with friends, family and colleagues, including 255.256.555.2 in contexts where a call would be when one knows the other person is busy 192.169.1.23 with family or work activities).; 172.01.123.254:1234\nfor example, to order products or 10.1.23.7 http://www_example.com/ 255.255.255.000 09.09.09.09\nservices fromhttps://regex-performance.github.io/exercises.html 3...3 0.0.0.0:22 0.0.0.0:68768\nthis permits communication even between busy individuals255.253.123.2:8000 https: // . Text messages can also http:// be used to http://example.com/path?name=Branch&products=[Journeys,Email,Universal%20Ads]interact with automated systems,https:// regex -performance.github.io/ exercises.html172.01.123.999:1234\nhttps//nan-academy.github.io/js-training/?page=editor#data.nested impolite or inappropriate (e.g., calling very late at night orhttp://localhost/exercises\nhttps://192.168.1.123?something=nothing&pro=[23] htts:/nan-academy.github.io/js-training?b=123&a=123/?page=editor#data.nested  Like e-mail and voicemail and unlike calls (in which the caller hopes to speak directly with the recipient),\nhttp://www.example.com/catalog.asp?itemid=232&template=fresh&crcat=ppc&crsource=google&crkw=buy-a-lot texting does not require the caller and recipient to both be free at the same moment0.0.0.0';
console.log(findIP(str))
