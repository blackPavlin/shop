/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { LoginRequest } from '../models/LoginRequest';
import type { LoginResponse } from '../models/LoginResponse';
import type { SignupRequest } from '../models/SignupRequest';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class AuthService {

    /**
     * Авторизация пользователя
     * @param requestBody 
     * @returns LoginResponse successful operation
     * @throws ApiError
     */
    public static postAuthLogin(
requestBody?: LoginRequest,
): CancelablePromise<LoginResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/login',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Регистрация пользователя
     * @param requestBody 
     * @returns LoginResponse successful operation
     * @throws ApiError
     */
    public static postAuthSignup(
requestBody?: SignupRequest,
): CancelablePromise<LoginResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/signup',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

}
