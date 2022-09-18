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
     * @returns User successful operation
     * @throws ApiError
     */
    public static getUser(): CancelablePromise<User> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/user',
        });
    }

}
