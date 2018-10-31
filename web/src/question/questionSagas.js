import { takeLatest, call, put } from 'redux-saga/effects'
import axios from 'axios'
import { types, actions } from './questionRedux'

export function * watcherSaga () {
  yield takeLatest(types.API_CALL_REQUEST, workerSaga)
}

function * workerSaga () {
  try {
    const response = yield call(fetchQuestion)
    const question = response.data

    yield put(actions.apiCallSuccess(question))
  } catch (error) {
    yield put(actions.apiCallFailure(error))
  }
}

function fetchQuestion () {
  return axios({
    method: 'post',
    url: 'https://api.dev.nettaton.com/question'
  })
}
