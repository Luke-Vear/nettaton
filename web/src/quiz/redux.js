export const types = {
  NEW_QUESTION_REQUEST: 'NEW_QUESTION_REQUEST',
  NEW_QUESTION_SUCCESS: 'NEW_QUESTION_SUCCESS',
  NEW_QUESTION_FAILURE: 'NEW_QUESTION_FAILURE',
  SEND_ANSWER_REQUEST: 'SEND_ANSWER_REQUEST',
  SEND_ANSWER_SUCCESS: 'SEND_ANSWER_SUCCESS',
  SEND_ANSWER_FAILURE: 'SEND_ANSWER_FAILURE',
  UPDATE_ANSWER: 'UPDATE_ANSWER'
}

export const actions = {
  newQuestionRequest: () => ({ type: types.NEW_QUESTION_REQUEST }),
  newQuestionSuccess: question => ({ type: types.NEW_QUESTION_SUCCESS, question }),
  newQuestionFailure: error => ({ type: types.NEW_QUESTION_FAILURE, error }),

  sendAnswerRequest: (question, answer) => ({ type: types.SEND_ANSWER_REQUEST, question, answer }),
  sendAnswerSuccess: result => ({ type: types.SEND_ANSWER_SUCCESS, result }),
  sendAnswerFailure: error => ({ type: types.SEND_ANSWER_FAILURE, error }),

  updateAnswer: answer => ({ type: types.UPDATE_ANSWER, answer })
}

const initialState = {
  busy: false,
  error: null,
  question: null,
  answer: '',
  result: null
}

export function reducer (state = initialState, action) {
  const logturn = (input) => {
    console.log(input)
    return input
  }
  switch (action.type) {
    case types.NEW_QUESTION_REQUEST:
      return logturn({ ...state, busy: true, error: null })
    case types.NEW_QUESTION_SUCCESS:
      return logturn({ ...state, busy: false, question: action.question })
    case types.NEW_QUESTION_FAILURE:
      return logturn({ ...state, busy: false, error: action.error })
    case types.SEND_ANSWER_REQUEST:
      return logturn({ ...state, busy: true, error: null })
    case types.SEND_ANSWER_SUCCESS:
      return logturn({ ...state, busy: false, question: null, result: action.result })
    case types.SEND_ANSWER_FAILURE:
      return logturn({ ...state, busy: false, error: action.error })
    case types.UPDATE_ANSWER:
      return logturn({ ...state, answer: action.answer })
    default:
      return logturn(state)
  }
}
