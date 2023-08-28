/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { User } from '../models/User';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class UserService {

    /**
     * Получить информацию о пользователе
     * @returns User OK
     * @throws ApiError
     */
    public static getUser(): CancelablePromise<User> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/user',
            errors: {
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

}
