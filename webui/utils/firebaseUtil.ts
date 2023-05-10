export enum FirebaseErrorCode {
    ACCOUNT_EXISTS = 'auth/account-exists-with-different-credential'
}

export const FIREBASE_ERROR_MESSAGE = {
    [FirebaseErrorCode.ACCOUNT_EXISTS]: 'Email already exists with other provider'
}