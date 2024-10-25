func generateParenthesis(n int) []string {
    var parenthesis []string
    var makeBalancedParenthesis func(openCount int, closeCount int, s string)

    makeBalancedParenthesis = func(openCount int, closeCount int, s string) {
        if openCount == closeCount && openCount == n {
            parenthesis = append(parenthesis, s)
            return
        }
        if openCount < n {
            addedOpenBracket := s + string('(')
            makeBalancedParenthesis(openCount+1, closeCount, addedOpenBracket)
        }
        if openCount > closeCount && closeCount < n {
            addedCloseBracket := s + string(')')
            makeBalancedParenthesis(openCount, closeCount+1, addedCloseBracket)
        }
    }

    makeBalancedParenthesis(0, 0, "")
    return parenthesis
}