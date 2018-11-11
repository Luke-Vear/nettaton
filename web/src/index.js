import React from 'react'
import ReactDOM from 'react-dom'

import '../public/styles.scss'
import App from './app/App'

import { Provider } from 'react-redux'
import { createStore, applyMiddleware, combineReducers } from 'redux'
import createSagaMiddleware from 'redux-saga'
import { all } from 'redux-saga/effects'

import { reducer as quizReducer } from './quiz/redux'
import { newQuestionWatcher, sendAnswerWatcher } from './quiz/sagas'

const sagaMiddleware = createSagaMiddleware()

const rootReducer = combineReducers({
  quiz: quizReducer
})

let store = createStore(rootReducer, applyMiddleware(sagaMiddleware))

function * rootSaga () {
  yield all([
    newQuestionWatcher(),
    sendAnswerWatcher()
  ])
}

sagaMiddleware.run(rootSaga)

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
)
