/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Category } from '../models/Category';
import type { CategoryList } from '../models/CategoryList';
import type { CreateCategoryRequest } from '../models/CreateCategoryRequest';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class CategoryService {

    /**
     * Получить список категорий товаров
     * @returns CategoryList OK
     * @throws ApiError
     */
    public static getCategory(): CancelablePromise<CategoryList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/category',
        });
    }

    /**
     * Создать категорию
     * @param requestBody
     * @returns Category OK
     * @throws ApiError
     */
    public static postCategory(
        requestBody?: CreateCategoryRequest,
    ): CancelablePromise<Category> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/category',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Изменить категорию
     * @param requestBody
     * @returns Category OK
     * @throws ApiError
     */
    public static patchCategory(
        requestBody?: Category,
    ): CancelablePromise<Category> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/category',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Удалить категорию
     * @param categoryId ID заказа
     * @returns void
     * @throws ApiError
     */
    public static deleteCategory(
        categoryId: number,
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/category/{categoryId}',
            path: {
                'categoryId': categoryId,
            },
        });
    }

}
