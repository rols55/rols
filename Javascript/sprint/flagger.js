const flags = (obj) => {
    let output = {
      alias: { h: 'help'},
      description: '',
    };
  
    const helpFlags = obj.help || Object.keys(obj);
    const numFlags = Object.values(helpFlags).length;
  
    for (let i = 0; i < numFlags; i++) {
        if (numFlags !== 0) {
             output.alias = {
                  h: 'help', i: 'invert', c: 'convert-map', a: 'assign'
              };
        }
      const flag = helpFlags[i];
      const alias = flag.slice(0, 1);
      const description = obj[flag];
      output.alias[alias] = flag;
      output.description += `-${alias}, --${flag}: ${description}`;
      if (i !== numFlags - 1) {
        output.description += '\n';
      }
    }
  
    return output;
  };
  
  
  //console.log(flags({}), { alias: { h: 'help' }, description: '' })
  console.log(flags({
    invert: 'inverts and object',
    'convert-map': 'converts the object to an array',
    assign: 'uses the function assign - assign to target object',
  }))
  console.log(flags({
    invert: 'inverts and object',
    'convert-map': 'converts the object to an array',
    assign: 'uses the function assign - assign to target object',
    help: ['assign', 'invert'],
  }))

  console.log(flags({
    invert: 'inverts and object',
    'convert-map': 'converts the object to an array',
    assign: 'uses the function assign - assign to target object',
    help: ['invert'],
  }))