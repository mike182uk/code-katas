const normalise = str =>
  str
    .toLowerCase()
    .match(/[a-z0-9]/g)

module.exports = str => {
  if (typeof str !== 'string') throw new Error('Supplied palindrome is not a string')

  str = normalise(str)

  for (let i = 0; i <= str.length; i++) {
    let char = str[i]
    let charToCompareIndex = str.length - i - 1

    if (char !== str[charToCompareIndex]) return false
  }

  return true
}
