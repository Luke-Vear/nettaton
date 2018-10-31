import React from 'react'
import { actions } from '../question/questionRedux'
import { connect } from 'react-redux'
import Question from '../question/Question'
import Display from '../display/Display'

export const App = () => (
  <div>
    <h1>NETTATON</h1>
    <Display />
    <Question />
  </div>
)

const mapStateToProps = state => ({
  fetching: state.fetching,
  question: state.question,
  error: state.error
})

const mapDispatchToProps = dispatch => ({
  onRequestDog: () => dispatch(actions.apiCallRequest())
})

export default connect(mapStateToProps, mapDispatchToProps)(App)
