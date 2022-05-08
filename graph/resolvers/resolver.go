package resolvers

import (
    "app-api/ent"

    "app-api/modules/custom_component"
    _shopModuleUseCase "app-api/modules/shop/usecase"
    _templateSectionModuleUseCase "app-api/modules/template_section/usecase"
    _themeModuleUseCase "app-api/modules/theme/usecase"
    _themeTemplateModuleUseCase "app-api/modules/theme_template/usecase"
    _todoModuleUseCase "app-api/modules/todo/usecase" // just an example
    _userModuleUseCase "app-api/modules/user/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
    Client                        *ent.Client
    TodoUseCase                   _todoModuleUseCase.TodoUseCase // just an example
    CustomComponentUseCase        custom_component.UseCase
    UserUseCase                   _userModuleUseCase.UserUseCase
    ShopUseCase                   _shopModuleUseCase.ShopUseCase
    TemplateSectionUseCase        _templateSectionModuleUseCase.TemplateSectionUseCase
    TemplateSectionVersionUseCase _templateSectionModuleUseCase.TemplateSectionVersionUseCase
    ThemeUseCase                  _themeModuleUseCase.ThemeUseCase
    ThemeTemplateUseCase          _themeTemplateModuleUseCase.ThemeTemplateUseCase
}
