import React from 'react'
import ReactDOM from 'react-dom'
import App from './app/App'

import { createStore, applyMiddleware, combineReducers } from 'redux'
import createSagaMiddleware from 'redux-saga'
import { Provider } from 'react-redux'

import { reducer as quizReducer } from './quiz/redux'
import { newQuestionWatcher, sendAnswerWatcher } from './quiz/sagas'

const sagaMiddleware = createSagaMiddleware()

const rootReducer = combineReducers({
  quiz: quizReducer
})

let store = createStore(rootReducer, applyMiddleware(sagaMiddleware))

sagaMiddleware.run(newQuestionWatcher, sendAnswerWatcher)

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
)
