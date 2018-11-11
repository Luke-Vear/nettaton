import { takeLatest, call, put } from 'redux-saga/effects'
import axios from 'axios'
import { types, actions } from './redux'

function buildEndpoint () {
  let env = 'f3dc7042-bc46-42d2-9f8f-41417d48ca4d'
  if (env.length === 36 || env.length === 0) {
    env = 'dev'
  }

  let ep = 'https://api.'
  if (env !== 'prod') {
    ep += env + '.'
  }
  ep += 'nettaton.com'

  return ep
}

const endpoint = buildEndpoint()

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
    url: endpoint + '/question'
  })
}

export function * sendAnswerWatcher () {
  yield takeLatest(types.SEND_ANSWER_REQUEST, sendAnswerSaga)
}

export function * sendAnswerSaga (action) {
  try {
    const response = yield call(sendAnswer, action.question, action.answer)
    const correct = response.data.correct

    yield put(actions.sendAnswerSuccess(correct))
  } catch (error) {
    yield put(actions.sendAnswerFailure(error))
  }
}

function sendAnswer (question, answer) {
  return axios({
    method: 'post',
    url: endpoint + '/question/' + question.id + '/answer',
    data: {
      answer: answer
    }
  })
}
