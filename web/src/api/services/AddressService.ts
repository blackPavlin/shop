/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Address } from '../models/Address';
import type { AddressList } from '../models/AddressList';
import type { CreateAddressRequest } from '../models/CreateAddressRequest';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class AddressService {

    /**
     * Получить список адресов доставки
     * @returns AddressList OK
     * @throws ApiError
     */
    public static getAddress(): CancelablePromise<AddressList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/address',
            errors: {
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Добавить адрес доставки
     * @param requestBody 
     * @returns Address OK
     * @throws ApiError
     */
    public static postAddress(
requestBody?: CreateAddressRequest,
): CancelablePromise<Address> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/address',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

}
