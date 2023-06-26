const neuron = (data) => {
    const output = {};
  
    for (const item of data) {
      const typeRegex = /^(.*?): (.*?) - Response: (.*)$/;
      const typeMatch = item.match(typeRegex);
  
      if (typeMatch) {
        const type = typeMatch[1].toLowerCase();
        const question = typeMatch[2];
        const response = typeMatch[3];
  
        const key = question.toLowerCase().replace(/\s+/g, '_').replace(/[^\w\s]/g, '');
  
        if (!output[type]) {
          output[type] = {};
        }
  
        if (!output[type][key]) {
          output[type][key] = { [type.slice(0, -1)]: question, responses: [] };
        }
  
        output[type][key].responses.push(response);
      }
    }
  
    return output;
  };
  