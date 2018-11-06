import { takeLatest, call, put } from 'redux-saga/effects'
import axios from 'axios'
import { types, actions } from './redux'

export function * newQuestionWatcher () {
  yield takeLatest(types.NEW_QUESTION_REQUEST, newQuestionSaga)
}

export function * newQuestionSaga () {
  try {
    const response = yield call(newQuestion)
    const question = response.data

    yield put(actions.newQuestionSuccess(question))
  } catch (error) {
    yield put(actions.newQuestionFailure(error))
  }
}

function newQuestion () {
  return axios({
    method: 'post',
    url: 'https://api.dev.nettaton.com/question'
  })
}

export function * sendAnswerWatcher () {
  yield takeLatest(types.SEND_ANSWER_REQUEST, sendAnswerSaga)
}

export function * sendAnswerSaga (action) {
  console.log('sendAnswerSaga:', action.question)
  try {
    const response = yield call(sendAnswer, action.question, action.answer)
    const result = response.data

    yield put(actions.sendAnswerSuccess(result))
  } catch (error) {
    yield put(actions.sendAnswerFailure(error))
  }
}

function sendAnswer (question, answer) {
  console.log('sendAnswer')
  return axios({
    method: 'post',
    url: 'https://api.dev.nettaton.com/question/' + question.id + '/answer',
    data: {
      answer: answer
    }
  })
}
