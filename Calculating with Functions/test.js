const {
    zero, one, two, three, four, five, six, seven, eight, nine,
    plus, minus, times, dividedBy,
} = require("./solution");
const {assert} = require("chai");


describe("Tests", () => {
    it("test", () => {
        assert.equal(seven(times(five())), 35);
        assert.equal(four(plus(nine())), 13);
        assert.equal(eight(minus(three())), 5);
        assert.equal(six(dividedBy(two())), 3);
    });
});
