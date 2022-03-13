const {assert} = require("chai");
const {toAndFrom} = require("./solution");

describe("To And From",()=>{ 
  it("example tests",()=>{
    assert.strictEqual(toAndFrom(53,55,32),53);
    assert.strictEqual(toAndFrom(66,17,49),17);
    assert.strictEqual(toAndFrom(2,4,3),3);
    assert.strictEqual(toAndFrom(1,10,8),9);
    assert.strictEqual(toAndFrom(10,4,8),6);
    assert.strictEqual(toAndFrom(2,4,5),3);
    assert.strictEqual(toAndFrom(42,150,548),142);
  });
});