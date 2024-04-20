// getURL function
function getURL(dataSet) {
    const urlRegex = /(https?:\/\/[^\s]+)/g;
    return dataSet.match(urlRegex) || [];
  }
  
  // greedyQuery function
  function greedyQuery(dataSet) {
    const urlRegex = /https?:\/\/[^\s]+/g;
    const urls = dataSet.match(urlRegex) || [];
    const filteredUrls = urls.filter(url => {
      const params = url.match(/(?:\?|&)([^=&]+=[^&]*)/g) || [];
      return params.length >= 3;
    });
    return filteredUrls;
  }
  
  // notSoGreedy function
  function notSoGreedy(dataSet) {
    const urls = getURL(dataSet);
    const result = [];
  
    for (const url of urls) {
      const urlObject = new URL(url);
      const queryParams = urlObject.searchParams;
      if (queryParams && queryParams.toString() && queryParams.toString().split('&').length >= 2 && queryParams.toString().split('&').length <= 3) {
        result.push(url);
      }
    }
  
    return result;
  }
  