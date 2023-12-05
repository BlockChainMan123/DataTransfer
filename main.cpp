#include <iostream>
#include <openssl/ec.h>
#include <openssl/obj_mac.h>
#include <openssl/ecdsa.h>
#include <openssl/err.h>
#include <openssl/sha.h>
#include <openssl/evp.h>
#include <openssl/rand.h>
#include <openssl/bn.h>
#include <cstdint>
#include <fstream>
#include<time.h>
void ISE_keygen(const EC_GROUP *group, EC_POINT *public_key, BIGNUM *private_key);
bool ISE_Verify(const EC_GROUP *group, EC_POINT *public_key, const char* data, BIGNUM *r, BIGNUM *s);
int ISE_Signature(const EC_GROUP *group, const EC_POINT *public_key,const BIGNUM *private_key, char *data, BIGNUM *r, BIGNUM *s);
void ISE_encrypt(const EC_GROUP *group, const EC_POINT *public_key, const char *message, EC_POINT *ciphertext1, EC_POINT *ciphertext2);
void ISE_decrypt(const EC_GROUP *group, const EC_POINT *public_key,BIGNUM *private_key , const EC_POINT *ciphertext1, const EC_POINT *ciphertext2, BIGNUM *decrypted_message);
# define NumberExperiment 1
# define repeat 1000
using namespace std;
int main() {

    EC_KEY *key = EC_KEY_new_by_curve_name(NID_secp256k1);
    EC_KEY_generate_key(key);
    const EC_GROUP *group = EC_KEY_get0_group(key);
    EC_POINT *pub_key = EC_POINT_new(group);
    BIGNUM *priv_key = BN_new() ;
    ISE_keygen(group,pub_key, priv_key);
    char *message = "hellow";
    BIGNUM *r = BN_new();
    BIGNUM *s = BN_new();
    if (ISE_Verify(group, pub_key,message, r, s)== 1)
    {
        printf("verification correct \n");
    }
    else{
        printf("verification failed \n");
    }
    char *messageaaa = "123456787654321";
    EC_POINT *C1 = EC_POINT_new(group);
    EC_POINT *C2 = EC_POINT_new(group);
    ISE_encrypt(group, pub_key, messageaaa, C1, C2);
    BIGNUM *messageaaa_dec = BN_new();
    ISE_decrypt(group, pub_key,priv_key, C1, C2, messageaaa_dec);
    double long begin[NumberExperiment]={0}, end[NumberExperiment]={0}, duration[NumberExperiment]={0};
    int32_t kk =0; int i=0;
    BIGNUM *decrypted_message = BN_new();
    const char *message_enc = "12345677777777779";
    std::ofstream outfile;
    outfile.open("performance.txt");
    outfile<<"\nThe time consumption is in ms \n\n"<<std::endl;
    outfile<<"index  \t" <<"keygen   \t" <<"Sign   \t" <<"Verify   \t"<<"encrypt   \t" <<"decrypt"<<std::endl;
    std::cout<<"index \t" <<"keygen  \t"<<"Sign    \t" <<"Verify   \t" <<"encrypt  \t " <<"decrypt"<<std::endl;
for(i=0; i<NumberExperiment; i++)
{
    begin[i] = clock(); kk= repeat;
    while(kk--)
        ISE_keygen(group,pub_key, priv_key);
    end[i] = clock();
    duration[i] = 1.0*(end[i]-begin[i])/CLOCKS_PER_SEC;
    outfile<<i << "\t" <<duration[i]<< "\t ";
    std::cout<<i<< "\t"<<duration[i]<<"\t";

    begin[i] = clock(); kk= repeat;
    while(kk--)
        ISE_Signature(group,pub_key,priv_key,message, r, s);
    end[i] = clock();
    duration[i] = 1.0*(end[i]-begin[i])/CLOCKS_PER_SEC;
    outfile<<duration[i]<< "\t ";
    std::cout<<duration[i]<<"\t";

    begin[i] = clock(); kk= repeat;
    while(kk--)
        ISE_Verify(group, pub_key,message, r, s);
    end[i] = clock();
    duration[i] = 1.0*(end[i]-begin[i])/CLOCKS_PER_SEC;
    outfile<<duration[i]<< "\t";
    std::cout<<duration[i]<<"\t ";

    begin[i] = clock(); kk= repeat;
    while(kk--)
        ISE_encrypt(group, pub_key, messageaaa, C1, C2);
    end[i] = clock();
    duration[i] = 1.0*(end[i]-begin[i])/CLOCKS_PER_SEC;
    outfile<<duration[i]<< " \t ";
    std::cout<<duration[i]<<"\t";

    begin[i] = clock(); kk= repeat;
    while(kk--)
    ISE_decrypt(group, pub_key,priv_key, C1, C2, messageaaa_dec);
    end[i] = clock();
    duration[i] = 1.0*(end[i]-begin[i])/CLOCKS_PER_SEC;
    outfile<<duration[i]<< std::endl;
    std::cout<<duration[i]<<std::endl;
}
return 0;
}
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
void ISE_keygen(const EC_GROUP *group, EC_POINT *public_key, BIGNUM *private_key)
{
    BIGNUM *k = BN_new();
    if (k ==NULL)
    {std::cout<<"faild to intialize k to a random number" <<std::endl;}
    BN_rand_range(k, EC_GROUP_get0_order(group));
    EC_POINT_mul(group,public_key,k,NULL,NULL,NULL);
    BN_copy(private_key, k);
}

void ISE_encrypt(const EC_GROUP *group, const EC_POINT *public_key, const char *message, EC_POINT *ciphertext1, EC_POINT *ciphertext2) 
{
    EC_POINT *kP = EC_POINT_new(group);
    BIGNUM *k = BN_new();
    BN_rand_range(k, EC_GROUP_get0_order(group));
    EC_POINT_mul(group, ciphertext1, k, NULL, NULL, NULL);
    EC_POINT_mul(group, kP, NULL, public_key, k, NULL);
    EC_POINT *message_2_EC_point = EC_POINT_new(group);
    BIGNUM *msg_bn = BN_new();
    BN_hex2bn(&msg_bn, message);
    EC_POINT_set_affine_coordinates_GFp(group, message_2_EC_point,msg_bn,msg_bn,NULL);
    EC_POINT_add(group, ciphertext2,kP, message_2_EC_point,NULL); 
    BN_free(msg_bn);
    BN_free(k);
    EC_POINT_free(message_2_EC_point);
    EC_POINT_free(kP);
}

void ISE_decrypt(const EC_GROUP *group, const EC_POINT *public_key,BIGNUM *private_key , const EC_POINT *ciphertext1, 
                const EC_POINT *ciphertext2, BIGNUM *decrypted_message)
{
    EC_POINT *invC1 = EC_POINT_new(group);
    invC1 = EC_POINT_dup(ciphertext1, group);
    EC_POINT_invert(group, invC1, NULL);
    EC_POINT *M = EC_POINT_new(group);
    EC_POINT *MM = EC_POINT_new(group);
    EC_POINT_mul(group, MM, NULL, invC1, private_key, NULL); 
    EC_POINT_add(group, MM, ciphertext2, MM,NULL);
    BIGNUM *y =  BN_new();
    EC_POINT_get_affine_coordinates(group, M,decrypted_message,y,NULL);
    // Free allocated memory
    BN_free(y);
    EC_POINT_free(MM);
    EC_POINT_free(invC1);
    EC_POINT_free(M);
}
int ISE_Signature(const EC_GROUP *group, const EC_POINT *public_key,const BIGNUM *private_key, const char *data, BIGNUM *r, BIGNUM *s)
{
    unsigned char digest[SHA256_DIGEST_LENGTH]={0};
    SHA256((const unsigned char*)data, strlen(data), digest);
    BIGNUM *digenst_Bn = BN_new();
    BN_bin2bn(digest,32,digenst_Bn);
    EC_POINT *kG = EC_POINT_new(group);
    BIGNUM *r_d = BN_new();
    BIGNUM *M_plus_rd = BN_new();
    BIGNUM *k = BN_new();
    BN_CTX *ctx = BN_CTX_new();
    BIGNUM *k_mul_inv = BN_new();
    if (k ==NULL | k_mul_inv==NULL|M_plus_rd==NULL |ctx==NULL |kG==NULL)
    {
        std::cout<<"faild to intialize k, k_inv, ctx,kG " <<std::endl;
    }
    BN_rand_range(k, EC_GROUP_get0_order(group));
    EC_POINT_mul(group, kG, k, NULL, NULL, NULL);
    EC_POINT_get_affine_coordinates(group, kG,r,r,NULL);
    BN_mod_inverse(k_mul_inv,k,EC_GROUP_get0_order(group),ctx);
    BN_mod_mul(r_d,r,private_key, EC_GROUP_get0_order(group),ctx);
    BN_mod_add(M_plus_rd,digenst_Bn,r_d,EC_GROUP_get0_order(group),ctx);
    BN_mod_mul(s, k_mul_inv,M_plus_rd,EC_GROUP_get0_order(group),ctx);
    if (s==NULL | r == NULL)
    {
        std::cerr << "Error Occured when generating ECDSA signature." << std::endl;
            BN_free(digenst_Bn);
            BN_free(k);
            BN_free(r_d);
            BN_free(k_mul_inv);
            BN_free(M_plus_rd);
            BN_CTX_free(ctx);
    }
    BN_free(digenst_Bn);
    BN_free(k);
    BN_free(r_d);
    BN_free(k_mul_inv);
    BN_free(M_plus_rd);
    BN_free(y);
    BN_CTX_free(ctx);
    
    return 1;
}
bool ISE_Verify(const EC_GROUP *group, EC_POINT *public_key, const char* data, BIGNUM *r, BIGNUM *s)
{
    unsigned char digest[SHA256_DIGEST_LENGTH]={0};
    SHA256((const unsigned char*)data, strlen(data), digest);
    BIGNUM *digenst_Bn = BN_new();
    BN_bin2bn(digest,32,digenst_Bn);
    BIGNUM *s_inv = BN_new();
    BIGNUM *m_s_inv = BN_new();
    BIGNUM *r_s_inv = BN_new();
    BIGNUM *r_x_coerdinate = BN_new();
    BN_CTX *ctx = BN_CTX_new();
    EC_POINT *m_s_inv_to_point = EC_POINT_new(group);
    EC_POINT *r_s_inv_to_point = EC_POINT_new(group);
    EC_POINT *r_prime = EC_POINT_new(group);
    BN_mod_inverse(s_inv,s,EC_GROUP_get0_order(group),NULL);  
    BN_mod_mul(m_s_inv,digenst_Bn,s_inv, EC_GROUP_get0_order(group),ctx); 
    BN_mod_mul(r_s_inv,r,s_inv, EC_GROUP_get0_order(group),ctx);          
    EC_POINT_mul(group, m_s_inv_to_point, m_s_inv, NULL, ctx, NULL); 
    EC_POINT_mul(group, r_s_inv_to_point, NULL, public_key, r_s_inv, NULL);               
    EC_POINT_add(group,r_prime,r_s_inv_to_point, m_s_inv_to_point,NULL); 
    EC_POINT_get_affine_coordinates(group, r_prime,r_x_coerdinate,r_x_coerdinate,NULL);
    if (BN_cmp(r_x_coerdinate,r) != 0)
    {
        printf("Verification Failed! \n");
        BN_free(digenst_Bn);
        BN_free(r_x_coerdinate);
        BN_free(y);
        BN_free(s_inv);
        BN_free(m_s_inv);
        BN_free(r_s_inv);
        BN_CTX_free(ctx);
        EC_POINT_free(r_prime);
        EC_POINT_free(m_s_inv_to_point);
        EC_POINT_free(r_s_inv_to_point);
        return false;

    }
    BN_free(digenst_Bn);
    BN_free(r_x_coerdinate);
    BN_free(y);
    BN_free(s_inv);
    BN_free(m_s_inv);
    BN_free(r_s_inv);
    BN_CTX_free(ctx);
    EC_POINT_free(r_prime);
    EC_POINT_free(m_s_inv_to_point);
    EC_POINT_free(r_s_inv_to_point);

    return true;
}
