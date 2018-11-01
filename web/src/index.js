import React from 'react'
import ReactDOM from 'react-dom'
import App from './app/App'

import { createStore, applyMiddleware } from 'redux'
import createSagaMiddleware from 'redux-saga'
import { Provider } from 'react-redux'

import { reducer } from './question/questionRedux'
import { watcherSaga } from './question/questionSagas'

const sagaMiddleware = createSagaMiddleware()

let store = createStore(reducer, applyMiddleware(sagaMiddleware))

sagaMiddleware.run(watcherSaga)

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
)
