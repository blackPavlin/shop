/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Cart } from '../models/Cart';
import type { CartList } from '../models/CartList';
import type { CartProduct } from '../models/CartProduct';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class CartService {

    /**
     * Получить корзину пользователя
     * @returns CartList OK
     * @throws ApiError
     */
    public static getCart(): CancelablePromise<CartList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/cart',
            errors: {
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Добавить товар в корзину
     * @param requestBody
     * @returns Cart OK
     * @throws ApiError
     */
    public static postCart(
        requestBody?: CartProduct,
    ): CancelablePromise<Cart> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/cart',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Изменить количество товара в корзине
     * @param requestBody
     * @returns Cart OK
     * @throws ApiError
     */
    public static patchCart(
        requestBody?: CartProduct,
    ): CancelablePromise<Cart> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/cart',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Очистить корзину
     * @returns void
     * @throws ApiError
     */
    public static deleteCart(): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/cart',
            errors: {
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Удалить товар из корзины
     * @param cartId ID товара в корзине
     * @returns void
     * @throws ApiError
     */
    public static deleteCart1(
        cartId: string,
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/cart/{cartId}',
            path: {
                'cartId': cartId,
            },
            errors: {
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

}
