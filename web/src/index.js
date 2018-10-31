import React from 'react'
import ReactDOM from 'react-dom'
import App from './app/App'

import { createStore, applyMiddleware } from 'redux'
import createSagaMiddleware from 'redux-saga'
import { Provider } from 'react-redux'

import { reducer as questionReducer } from './question/questionRedux'
import { watcherSaga as questionWatcherSaga } from './question/questionSagas'

const sagaMiddleware = createSagaMiddleware()

let store = createStore(questionReducer, applyMiddleware(sagaMiddleware))

sagaMiddleware.run(questionWatcherSaga)

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
)
