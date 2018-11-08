export const types = {
  NEW_QUESTION_REQUEST: 'NEW_QUESTION_REQUEST',
  NEW_QUESTION_SUCCESS: 'NEW_QUESTION_SUCCESS',
  NEW_QUESTION_FAILURE: 'NEW_QUESTION_FAILURE',
  SEND_ANSWER_REQUEST: 'SEND_ANSWER_REQUEST',
  SEND_ANSWER_SUCCESS: 'SEND_ANSWER_SUCCESS',
  SEND_ANSWER_FAILURE: 'SEND_ANSWER_FAILURE',
  ANSWER_UPDATE: 'ANSWER_UPDATE'
}

export const actions = {
  newQuestionRequest: () => ({ type: types.NEW_QUESTION_REQUEST }),
  newQuestionSuccess: question => ({ type: types.NEW_QUESTION_SUCCESS, question }),
  newQuestionFailure: error => ({ type: types.NEW_QUESTION_FAILURE, error }),

  sendAnswerRequest: (question, answer) => ({ type: types.SEND_ANSWER_REQUEST, question, answer }),
  sendAnswerSuccess: correct => ({ type: types.SEND_ANSWER_SUCCESS, correct }),
  sendAnswerFailure: error => ({ type: types.SEND_ANSWER_FAILURE, error }),

  answerUpdate: answer => ({ type: types.ANSWER_UPDATE, answer })
}

const initialState = {
  busy: false,
  error: null,
  question: null,
  answer: '',
  correct: null
}

export function reducer (state = initialState, action) {
  switch (action.type) {
    case types.NEW_QUESTION_REQUEST:
      return { ...state, busy: true, error: null }
    case types.NEW_QUESTION_SUCCESS:
      return { ...state, busy: false, question: action.question, correct: null }
    case types.NEW_QUESTION_FAILURE:
      return { ...state, busy: false, error: action.error }
    case types.SEND_ANSWER_REQUEST:
      return { ...state, busy: true, error: null }
    case types.SEND_ANSWER_SUCCESS:
      return action.correct === true ? { ...state, busy: false, correct: action.correct, question: null } : { ...state, busy: false, correct: action.correct }
    case types.SEND_ANSWER_FAILURE:
      return { ...state, busy: false, error: action.error }
    case types.ANSWER_UPDATE:
      return { ...state, answer: action.answer, error: null }
    default:
      return state
  }
}
