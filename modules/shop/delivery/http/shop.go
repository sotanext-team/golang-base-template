package http

import (
	"strconv"
	"time"

	"golang-base/models"
	"golang-base/modules/shop/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ShopHandler struct {
	shopUseCase usecase.ShopUseCase
}

func NewShopHandler(db *gorm.DB) ShopHandler {
	shopUseCase := usecase.NewShopUseCase(db)
	return ShopHandler{
		shopUseCase: shopUseCase,
	}
}

func (instance *ShopHandler) ShopRegister(c *gin.Context) {
	var shop models.Shop
	err := c.ShouldBindJSON(&shop)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	shopResult, err := instance.shopUseCase.CreateShop(c.Request.Context(), shop)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": shopResult})
}

func (instance *ShopHandler) List(c *gin.Context) {
	shopResult, err := instance.shopUseCase.GetShops(c.Request.Context())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": shopResult})
}

func (instance *ShopHandler) Get(c *gin.Context) {
	shopID := c.Param("shop_id")
	number, err := strconv.ParseUint(shopID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	shopResult, err := instance.shopUseCase.GetShop(c.Request.Context(), number)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": shopResult})
}

func (instance *ShopHandler) TestPermission(c *gin.Context) {

	// userValue, _ := c.Get("currentShop")
	// currentShop := userValue.(models.Shop)

	// sub := strconv.Itoa(int(currentUser.ID))
	// e.AddGroupingPolicy()

	// logrus.Info(e.AddPolicy("alice_admin", "data_alice_group", "write"))
	// logrus.Info(e.AddPolicy("alice_admin", "data_alice_group", "read"))
	// logrus.Info(e.AddPolicy("alice_admin", "data_alice_group", "delete"))
	// logrus.Info(e.AddPolicy("alice_admin", "data_alice_section", "write"))
	// logrus.Info(e.AddPolicy("alice_admin", "data_alice_section", "read"))
	// logrus.Info(e.AddPolicy("alice_admin", "data_alice_section", "delete"))

	// logrus.Info(e.AddPolicy("bob_admin", "data_bob_group", "write"))
	// logrus.Info(e.AddPolicy("bob_admin", "data_bob_group", "read"))
	// logrus.Info(e.AddPolicy("bob_admin", "data_bob_group", "delete"))
	// logrus.Info(e.AddPolicy("bob_admin", "data_bob_section", "write"))
	// logrus.Info(e.AddPolicy("bob_admin", "data_bob_section", "read"))
	// logrus.Info(e.AddPolicy("bob_admin", "data_bob_section", "delete"))

	// logrus.Info(e.AddPolicy("bob_manager", "data_bob_group", "write"))
	// logrus.Info(e.AddPolicy("bob_manager", "data_bob_group", "read"))
	// logrus.Info(e.AddPolicy("bob_manager", "data_bob_section", "write"))
	// logrus.Info(e.AddPolicy("bob_manager", "data_bob_section", "read"))

	// logrus.Info(e.AddNamedGroupingPolicy("g", "alice", "alice_admin"))
	// logrus.Info(e.AddNamedGroupingPolicy("g", "bob", "bob_admin"))
	// logrus.Info(e.AddNamedGroupingPolicy("g", "alice", "bob_manager"))
	// logrus.Info(e.AddNamedGroupingPolicy("g2", "alice_shop_1", "data_alice_group"))
	// logrus.Info(e.AddNamedGroupingPolicy("g2", "alice_shop_2", "data_alice_group"))
	// logrus.Info(e.AddNamedGroupingPolicy("g2", "alice_section_1", "data_alice_section"))
	// logrus.Info(e.AddNamedGroupingPolicy("g2", "alice_section_2", "data_alice_section"))

	for i := 0; i <= 10; i++ {
		logrus.Info(i)
		// user role
		// resource role

		// p, owner, shop_1, product, *
		// p, owner, shop_1, section, *
		// p, owner, shop_1, shop, *

		// p, shop_read, shop_1, shop, read
		// p, shop_write, shop_1, shop, write
		// p, product_read, shop_1, product, read
		// p, product_write, shop_1, product, write
		// p, product_delete, shop_1, product, delete
		// p, section_read, shop_1, section, read
		// p, section_write, shop_1, section, write
		// p, section_delete, shop_1, section, delete

		// g, user_1, owner, shop_1
		// g, admin, shop_read, shop_1
		// g, admin, shop_write, shop_1
		// g, admin, product_read, shop_1
		// g, admin, product_write, shop_1
		// g, admin, product_delete, shop_1
		// g, admin, section_read, shop_1
		// g, admin, section_write, shop_1
		// g, admin, section_delete, shop_1

		// add role no log
		// e.AddPolicy("owner", domain, "*")

		// e.AddPolicy("shop_read", domain, "shop_read")
		// e.AddPolicy("shop_write", domain, "shop_write")
		// e.AddPolicy("product_read", domain, "product_read")
		// e.AddPolicy("product_write", domain, "product_write")
		// e.AddPolicy("product_delete", domain, "product_delete")
		// e.AddPolicy("section_read", domain, "section_read")
		// e.AddPolicy("section_write", domain, "section_write")
		// e.AddPolicy("section_delete", domain, "section_delete")

		// authz.AddRoleToDomain(uint(i), uint(i), authz.OWNER_ROLE)
		// authz.AddRoleToDomain(uint(i), uint(i), authz.SHOP_READ)
		// authz.AddRoleToDomain(uint(i), uint(i), authz.SHOP_WRITE)
		// authz.AddRoleToDomain(uint(i), uint(i), authz.PRODUCT_READ)
		// authz.AddRoleToDomain(uint(i), uint(i), authz.PRODUCT_WRITE)
		// authz.AddRoleToDomain(uint(i), uint(i), authz.PRODUCT_DELETE)
		// authz.AddRoleToDomain(uint(i), uint(i), authz.SECTION_READ)
		// authz.AddRoleToDomain(uint(i), uint(i), authz.SECTION_WRITE)
		// authz.AddRoleToDomain(uint(i), uint(i), authz.SECTION_DELETE)

		// for j := 1; j < 50; j++ {
		// 	authz.AddRoleToDomain(uint(i), uint(j), authz.ADMIN_ROLE)
		// }
	}

	// test speed
	t1 := time.Now()

	// logrus.Info(e.Enforce("admin", "shop_1", "product", "read"))
	// logrus.Info(e.Enforce("user_1", "shop_1", "product", "write"))
	// logrus.Info(e.Enforce("user_1", "shop_1", "product", "delete"))

	// logrus.Info(e.Enforce("user_1", "shop_2", "product", "delete"))
	// logrus.Info(e.Enforce("user_1", "shop_2", "product", "read"))
	// logrus.Info(e.Enforce("user_1", "shop_2", "shop", "read"))
	// logrus.Info(e.Enforce("user_1", "shop_2", "shop", "write"))
	// logrus.Info(e.Enforce("user_1", "shop_2", "shop", "delete"))

	logrus.Info(time.Since(t1))
	//

	// logrus.Info(e.Enforce("alice", "alice_shop_1", "read"))
	// logrus.Info(e.Enforce("alice", "alice_shop_2", "delete"))
	// logrus.Info(e.Enforce("alice", "alice_section_1", "read"))
	// logrus.Info(e.Enforce("alice", "alice_section_2", "delete"))
	// logrus.Info(e.Enforce("alice", "data_bob_section", "write"))

	// permission find

	// logrus.Info(e.GetAllRoles())
	// logrus.Info(e.GetAllActions())
	// logrus.Info(e.GetRolesForUser("alice"))
	// logrus.Info(e.GetImplicitRolesForUser("alice"))
	// logrus.Info(e.GetImplicitPermissionsForUser("alice"))
	// logrus.Info(e.GetGroupingPolicy())

	// logrus.Info(e.DeleteRoleForUser("alice", "bob_manager"))
	// e.RemoveGroupingPolicy("g", "alice", "bob_manager")
	// logrus.Info(e.Enforce("alice", "data_bob_section", "write"))
}

func (instance *ShopHandler) TestCheckPermission(c *gin.Context) {
	// test speed
	t1 := time.Now()

	//

	// logrus.Info(e.Enforce("alice", "alice_shop_1", "read"))
	// logrus.Info(e.Enforce("alice", "alice_shop_2", "delete"))
	// logrus.Info(e.Enforce("alice", "alice_section_1", "read"))
	// logrus.Info(e.Enforce("alice", "alice_section_2", "delete"))
	// logrus.Info(e.Enforce("alice", "data_bob_section", "write"))

	// permission find

	// logrus.Info(e.GetAllRoles())
	// logrus.Info(e.GetAllActions())
	// logrus.Info(e.GetImplicitRolesForUser("alice"))
	// logrus.Info(e.GetImplicitPermissionsForUser("alice"))
	// logrus.Info(e.GetGroupingPolicy())

	// logrus.Info(e.DeleteRoleForUser("alice", "bob_manager"))
	// e.RemoveGroupingPolicy("g", "alice", "bob_manager")
	// logrus.Info(e.Enforce("alice", "data_bob_section", "write"))
	for i := 0; i < 1; i++ {
		go func() {

			t := time.Now()
			// logrus.Info(e.GetImplicitRolesForUser("owner", "shop_1"))
			// logrus.Info(e.GetImplicitRolesForUser("user_1", "shop_1"))
			// logrus.Info(e.GetImplicitRolesForUser("admin", "shop_1"))
			// logrus.Info(e.GetImplicitRolesForUser("user_1", "shop_2"))
			// logrus.Info(e.GetImplicitPermissionsForUser("user_1", "shop_2"))
			// logrus.Info(e.Enforce("owner", "shop_1", "product_read"))
			logrus.Info(time.Since(t))
			// t = time.Now()
			// logrus.Info(e.Enforce("user_1", "shop_1", "product", "write"))
			// logrus.Info(time.Since(t))
			// t = time.Now()
			// logrus.Info(e.Enforce("user_1", "shop_1", "product", "delete"))
			// logrus.Info(time.Since(t))
			// t = time.Now()
			// logrus.Info(e.Enforce("user_1", "shop_2", "product", "delete"))
			// logrus.Info(time.Since(t))
			// t = time.Now()
			// logrus.Info(e.Enforce("user_1", "shop_2", "product", "read"))
			// logrus.Info(time.Since(t))
			// t = time.Now()
			// logrus.Info(e.Enforce("user_1", "shop_2", "shop", "read"))
			// logrus.Info(time.Since(t))
			// t = time.Now()
			// logrus.Info(e.Enforce("user_1", "shop_2", "shop", "write"))
			// logrus.Info(time.Since(t))
			// t = time.Now()
			// logrus.Info(e.Enforce("user_1", "shop_2", "shop", "delete"))
			// logrus.Info(time.Since(t))
			// t = time.Now()
		}()
	}

	logrus.Info(time.Since(t1))
	//
}

func (instance *ShopHandler) FakeDB(c *gin.Context) {

}
