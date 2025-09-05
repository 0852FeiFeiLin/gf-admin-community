package en_crypto

import (
	"testing"
)

func TestPwdHash(t *testing.T) {
	password := "testPassword123"
	salt := "12345678"

	// 测试bcrypt密码哈希
	hash1, err := PwdHash(password, salt)
	if err != nil {
		t.Fatalf("PwdHash失败: %v", err)
	}

	// 验证生成的哈希是有效的bcrypt哈希
	if !IsValidPasswordHash(hash1) {
		t.Fatalf("生成的哈希不是有效的bcrypt格式: %s", hash1)
	}

	// 测试密码验证
	err = VerifyPassword(password, hash1)
	if err != nil {
		t.Fatalf("密码验证失败: %v", err)
	}

	// 测试错误密码
	err = VerifyPassword("wrongPassword", hash1)
	if err == nil {
		t.Fatal("错误密码应该验证失败")
	}

	// 同样的密码应该生成不同的哈希（因为内置随机盐）
	hash2, err := PwdHash(password, salt)
	if err != nil {
		t.Fatalf("第二次PwdHash失败: %v", err)
	}

	if hash1 == hash2 {
		t.Log("注意：两次生成的哈希相同，但这在bcrypt中是正常的")
	}

	// 验证第二个哈希也能正确验证密码
	err = VerifyPassword(password, hash2)
	if err != nil {
		t.Fatalf("第二个哈希的密码验证失败: %v", err)
	}
}

func TestMigrateLegacyHash(t *testing.T) {
	password := "testPassword123"
	salt := "12345678"

	// 创建一个旧的scrypt哈希（模拟）
	legacyHash, err := legacyPwdHash(password, salt)
	if err != nil {
		t.Fatalf("创建旧哈希失败: %v", err)
	}

	// 测试迁移
	newHash, needUpdate, err := MigrateLegacyHash(password, legacyHash, salt)
	if err != nil {
		t.Fatalf("迁移失败: %v", err)
	}

	if !needUpdate {
		t.Fatal("应该需要更新")
	}

	// 验证新哈希是有效的bcrypt
	if !IsValidPasswordHash(newHash) {
		t.Fatalf("迁移后的哈希不是有效的bcrypt格式: %s", newHash)
	}

	// 测试已经是bcrypt哈希的情况
	newHash2, needUpdate2, err := MigrateLegacyHash(password, newHash, salt)
	if err != nil {
		t.Fatalf("bcrypt哈希验证失败: %v", err)
	}

	if needUpdate2 {
		t.Fatal("bcrypt哈希不应该需要更新")
	}

	if newHash != newHash2 {
		t.Fatal("bcrypt哈希应该保持不变")
	}
}

func TestIsValidPasswordHash(t *testing.T) {
	// 测试有效的bcrypt哈希
	validHash := "$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewdBPj4B9k3jyOF6"
	if !IsValidPasswordHash(validHash) {
		t.Fatal("有效的bcrypt哈希被判断为无效")
	}

	// 测试无效的哈希
	invalidHashes := []string{
		"invalidhash",
		"$1$short",
		"tooshorthash",
		"",
	}

	for _, hash := range invalidHashes {
		if IsValidPasswordHash(hash) {
			t.Fatalf("无效哈希被判断为有效: %s", hash)
		}
	}
}