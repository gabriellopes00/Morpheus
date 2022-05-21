import { MigrationInterface, QueryRunner, Table } from 'typeorm'

export class CreateAccountMigration implements MigrationInterface {
  name = this.constructor.name.concat('1653147077411')

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.createTable(
      new Table({
        name: 'test_table',
        columns: [
          { name: 'id', type: 'uuid', isPrimary: true },
          { name: 'referencedId', type: 'uuid', isNullable: false },
          { name: 'created_at', type: 'timestamp', default: 'now()' }
        ]
      }),
      true
    )
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.dropTable('test_table', true)
  }
}
