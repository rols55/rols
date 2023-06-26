const escapeStr = '`\\ / "\''
const arr = Object.freeze([4, '2'])
const obj = Object.freeze({
    str: 'example string',
    num: 42,
    bool: true,
    undef: undefined
  })
const nested = Object.freeze({
arr: Object.freeze([4, undefined, '2']),
obj: Object.freeze({
    str: 'example string',
    num: 42,
    bool: true
    })
})