package en_crypto

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
	"github.com/gogf/gf/v2/text/gstr"
)

// PwdHash 使用bcrypt进行密码哈希加密
// bcrypt内置盐值管理，无需手动处理盐值
func PwdHash(text string, salt string) (string, error) {
	// 使用cost=12，在安全性和性能之间取得平衡
	// 对于2025年来说，cost=12是推荐的安全级别
	hash, err := bcrypt.GenerateFromPassword([]byte(text), 12)
	if err != nil {
		return "", fmt.Errorf("密码哈希生成失败: %w", err)
	}
	
	// bcrypt返回的hash已经包含了算法标识、cost、盐值和哈希值
	// 格式: $2a$12$saltsaltsaltsaltsaltsOuHashHashHashHashHashHashHash
	return string(hash), nil
}

// VerifyPassword 验证密码是否匹配
// 用于验证用户输入的密码与存储的哈希值是否匹配
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// IsValidPasswordHash 检查是否为有效的bcrypt哈希
func IsValidPasswordHash(hash string) bool {
	// bcrypt哈希长度通常是60字符，以$2开头
	return len(hash) == 60 && (hash[:3] == "$2a" || hash[:3] == "$2b" || hash[:3] == "$2y")
}

// MigrateLegacyHash 迁移旧的scrypt哈希到bcrypt
// 在用户下次登录时验证旧密码并更新为新哈希
func MigrateLegacyHash(password, legacyHash, salt string) (newHash string, needUpdate bool, err error) {
	// 如果已经是bcrypt哈希，直接验证
	if IsValidPasswordHash(legacyHash) {
		err = VerifyPassword(password, legacyHash)
		if err != nil {
			return "", false, err
		}
		return legacyHash, false, nil
	}
	
	// 验证旧的scrypt哈希
	oldHash, err := legacyPwdHash(password, salt)
	if err != nil {
		return "", false, fmt.Errorf("旧密码验证失败: %w", err)
	}
	
	// 比较哈希值
	if oldHash != legacyHash {
		return "", false, fmt.Errorf("密码验证失败")
	}
	
	// 生成新的bcrypt哈希
	newHash, err = PwdHash(password, salt)
	if err != nil {
		return "", false, fmt.Errorf("新密码哈希生成失败: %w", err)
	}
	
	return newHash, true, nil
}

// legacyPwdHash 保留旧的scrypt实现用于迁移验证
func legacyPwdHash(text string, salt string) (string, error) {
	// 复制原来的scrypt实现逻辑
	saltLen := len(salt)
	for saltLen < 8 {
		salt += "0"
		saltLen++
	}

	salt = gstr.SubStr(salt, saltLen-8, 8)

	key, err := scrypt.Key([]byte(text), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		return "", fmt.Errorf("scrypt密码哈希失败: %w", err)
	}
	return fmt.Sprintf("%x", key), nil
}
